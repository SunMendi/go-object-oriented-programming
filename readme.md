# OOP with Go (`oop_with_go`)

A professional mini-project demonstrating **Object-Oriented Programming (OOP) concepts in Go**.  
This project illustrates **Encapsulation, Abstraction, Polymorphism, and Composition/Embedding** using a simple online store example with products, cart, and payment providers.

---

## 🎯 Key Concepts of OOP in Go

Go is not a classical OOP language like Java or C++. Instead, it provides OOP features through **structs, interfaces, and methods**. The design philosophy emphasizes **simplicity, flexibility, and composition over inheritance**.

---

## 1️⃣ Encapsulation

- Encapsulation is about **hiding data and protecting it**, allowing controlled access through methods.
- **Private fields** are not directly accessible outside the package.

**Example: Product stock management**
```go
type Product struct {
    name  string // private
    price float64
    stock int
}

func (p *Product) ReduceStock(amount int) bool { ... }
```
Only the system (methods) can modify private fields. Users cannot directly update sensitive data.

**Analogy:** In banking, only validated methods can modify account balance.

---

## 2️⃣ Abstraction

Abstraction hides implementation details and exposes only what’s necessary.

In Go, interfaces provide abstraction.

**Example: Payment system interface**
```go
type PaymentProvider interface {
    Pay(amount float64) error
}
```
Any payment type (`bKash`, `Nagad`, `Stripe`) can implement this interface. Cart only knows it can call `Pay()` without caring how it works internally.

---

## 3️⃣ Inheritance / Embedding (Composition)

Go does not support classical inheritance. Instead, it uses embedding (composition).

Embedding means one struct includes another, so it can use its fields and methods.

**Example: Person → Employee**
```go
type Person struct { 
    Name string
    Age int 
}
type Employee struct {
    Person  // embedding
    JobTitle string
    Salary   int
}
```

**Analogy:**
- Java/Python: Dog *is-an* Animal → classical inheritance
- Go: Dog *has-an* Animal → embedding/composition

```go
type Animal struct { Name string }
func (a *Animal) Speak() { fmt.Println("???") }

type Dog struct {
    Animal // Dog has an Animal
    Breed string
}

d := Dog{Animal: Animal{Name:"Rex"}, Breed:"Shepherd"}
d.Speak()     // "???"
fmt.Println(d.Name) // "Rex"
```

---

## 4️⃣ Polymorphism

Polymorphism = one interface, many implementations.

**Example: PaymentProvider interface**
```go
type Bkash struct{}
func (b Bkash) Pay(amount float64) error { ... }

type Nagad struct{}
func (n Nagad) Pay(amount float64) error { ... }
```
A single function can work with any struct that implements the interface.

```go
cart := NewCart(Bkash{})
cart.Checkout()  // works with any PaymentProvider
```

---

## 5️⃣ Go’s Philosophy

- No heavy class + inheritance system (like Java/C++).
- Types (structs) + functions + interfaces provide flexibility.
- Object-like behavior achieved using receiver methods.
- Functions can be standalone or tied to a struct when needed.
- Simplicity and clarity are prioritized over strict OOP enforcement.

---

## 🏗 Project Structure

```
oop_with_go/
├── main.go                 # Entry point
├── models/                 # Data models
│   └── product.go
├── repo/                   # Repositories for data persistence
│   ├── base_repo.go
│   └── product_repo.go
├── service/                # Business logic / services
│   ├── cart_service.go
│   └── payment_service.go
└── payment/                # Payment providers
    ├── payment_provider.go
    ├── bkash.go
    └── stripe.go
```

---

## 6️⃣ How It Works

- **Models:** Define core data structures like Product.
- **Repositories:** Handle data persistence (simulated via BaseRepo and ProductRepo).
- **Services:** Business logic like Cart operations, adding products, calculating totals, and checkout.
- **Payment Providers:** Implement PaymentProvider interface (Bkash, Stripe), demonstrating polymorphism.
- **Main:** Ties everything together.

**Example flow:**  
Create products → add to cart → checkout → pay using any provider.

---

## 7️⃣ Example Usage

```go
shirt := *models.NewProduct(1, "Shirt", 500, 10)
pant := *models.NewProduct(2, "Pant", 800, 5)

bkash := payment.Bkash{}
cart := service.NewCart(bkash)
cart.AddItem(shirt)
cart.AddItem(pant)
cart.Checkout()
```
- Products are added to cart if in stock (**encapsulation**).
- Checkout works with any payment provider (**polymorphism & abstraction**).

---

## ✅ Key Takeaways

- **Encapsulation:** Protect sensitive data in structs.
- **Abstraction:** Expose only the methods needed via interfaces.
- **Composition/Embedding:** Reuse code without inheritance.
- **Polymorphism:** One interface, multiple implementations.
- **Go philosophy:** Simplicity + clarity + flexibility.

---

## 👤 Author

Written by Mehedi Mahbub

---

## 📚 Project Purpose

Learn and demonstrate OOP concepts