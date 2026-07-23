package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID       int     `json:"id"`
	Customer string  `json:"customer"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
}

type Processor struct {
	mu         sync.Mutex
	Processed  int
	TotalValue float64
}

func (o Order) String() string {
	return fmt.Sprintf("Order #%d (%s): $%.2f [%s]", o.ID, o.Customer, o.Amount, o.Status)
}

func (p *Processor) Record(o Order) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Processed++
	p.TotalValue += o.Amount
}

type Validator interface {
	Validate(o Order) error
}

type AmountValidator struct {
	MaxAmount float64
}

func (v AmountValidator) Validate(o Order) error {
	if o.Amount <= 0 {
		return fmt.Errorf("order #%d has invalid amount: %.2f", o.ID, o.Amount)
	}
	if o.Amount > v.MaxAmount {
		return fmt.Errorf("order #%d exceeds max allowed amount (%.2f > %.2f)", o.ID, o.Amount, v.MaxAmount)
	}
	return nil
}

func makeStatusUpdater() func(o Order) Order {
	statuses := []string{"validated", "processed", "shipped"}
	step := 0
	return func(o Order) Order {
		if step < len(statuses) {
			o.Status = statuses[step]
			step++
		}
		return o
	}
}

func worker(id int, jobs <-chan Order, results chan<- string, validator Validator, processor *Processor, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range jobs {
		if err := validator.Validate(order); err != nil {
			results <- fmt.Sprintf("[Worker %d] REJECTED: %v", id, err)
			continue
		}
		processor.Record(order)
		results <- fmt.Sprintf("[Worker %d] OK: %s", id, order.String())
	}
}

func main() {
	orders := []Order{
		{ID: 1, Customer: "Kelly", Amount: 250.00},
		{ID: 2, Customer: "Amina", Amount: -50.00},
		{ID: 3, Customer: "Otieno", Amount: 9999.00},
		{ID: 4, Customer: "Wanjiru", Amount: 120.50},
		{ID: 5, Customer: "Brian", Amount: 75.25},
	}

	updateStatus := makeStatusUpdater()
	for i, o := range orders {
		orders[i] = updateStatus(o)
	}

	jobs := make(chan Order, len(orders))
	results := make(chan string, len(orders))
	processor := &Processor{}
	validator := AmountValidator{MaxAmount: 1000}

	var wg sync.WaitGroup
	numWorkers := 3

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, validator, processor, &wg)
	}

	for _, o := range orders {
		jobs <- o
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("---- Processing Orders ----")
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("\n---- Summary ----")
	fmt.Printf("Orders processed successfully: %d\n", processor.Processed)
	fmt.Printf("Total value processed: $%.2f\n", processor.TotalValue)

	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		fmt.Println("JSON error:", err)
		return
	}
	fmt.Println("\n---- Orders as JSON ----")
	fmt.Println(string(data))

	time.Sleep(50 * time.Millisecond)
}