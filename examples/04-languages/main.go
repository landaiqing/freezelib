package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("💻 Programming Languages Examples")
	fmt.Println("=================================")

	// Create output directory
	os.MkdirAll("output", 0755)

	// Run language examples
	popularLanguagesExample()
	languageComparisonExample()
	multiLanguageProjectExample()
	languageSpecificFeaturesExample()

	fmt.Println("\n✅ Language examples completed!")
	fmt.Println("📁 Check the 'output' directory for generated files.")
}

// Popular programming languages showcase
func popularLanguagesExample() {
	fmt.Println("\n🌟 Popular Languages Showcase")
	fmt.Println("-----------------------------")

	// Language examples with sample code
	languages := []struct {
		name string
		code string
		lang string
	}{
		{
			"go",
			`package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from Go!")
	
	// Goroutine example
	go func() {
		fmt.Println("This runs concurrently")
	}()
	
	time.Sleep(time.Millisecond * 100)
}`,
			"go",
		},
		{
			"python",
			`import asyncio
from dataclasses import dataclass

@dataclass
class User:
    name: str
    age: int
    email: str

async def fetch_user(user_id: int) -> User:
    # Simulate API call
    await asyncio.sleep(1)
    return User(name="John Doe", age=30, email="john@example.com")

async def main():
    user = await fetch_user(123)
    print(f"User: {user.name}, {user.age}, {user.email}")

if __name__ == "__main__":
    asyncio.run(main())`,
			"python",
		},
		{
			"javascript",
			`// Modern JavaScript example
const fetchData = async () => {
  try {
    const response = await fetch('https://api.example.com/data');
    const data = await response.json();
    
    // Destructuring and spread operator
    const { items, meta } = data;
    const allItems = [...items, { id: 'new', value: 42 }];
    
    // Array methods
    const filtered = allItems.filter(item => item.value > 10);
    
    return filtered;
  } catch (error) {
    console.error('Error fetching data:', error);
    return [];
  }
};

// Call the function
fetchData().then(result => console.log(result));`,
			"javascript",
		},
		{
			"rust",
			`use std::collections::HashMap;

#[derive(Debug)]
struct User {
    name: String,
    age: u32,
    active: bool,
}

fn main() {
    // Create a new User
    let user = User {
        name: String::from("Alice"),
        age: 28,
        active: true,
    };
    
    // Using a HashMap
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    
    // Pattern matching
    match user.age {
        0..=17 => println!("{} is underage", user.name),
        18..=64 => println!("{} is an adult", user.name),
        _ => println!("{} is a senior", user.name),
    }
}`,
			"rust",
		},
		{
			"java",
			`import java.util.List;
import java.util.stream.Collectors;

public class StreamExample {
    public static void main(String[] args) {
        // Create a list of names
        List<String> names = List.of("Alice", "Bob", "Charlie", "David");
        
        // Use streams to filter and transform
        List<String> filteredNames = names.stream()
            .filter(name -> name.length() > 4)
            .map(String::toUpperCase)
            .sorted()
            .collect(Collectors.toList());
        
        // Print the results
        System.out.println("Original names: " + names);
        System.out.println("Filtered names: " + filteredNames);
    }
}`,
			"java",
		},
		{
			"csharp",
			`using System;
using System.Collections.Generic;
using System.Linq;

namespace LinqExample
{
    class Program
    {
        static void Main(string[] args)
        {
            // Create a list of products
            var products = new List<Product>
            {
                new Product { Id = 1, Name = "Laptop", Price = 1200.00m, Category = "Electronics" },
                new Product { Id = 2, Name = "Desk Chair", Price = 250.50m, Category = "Furniture" },
                new Product { Id = 3, Name = "Coffee Maker", Price = 89.99m, Category = "Kitchen" },
                new Product { Id = 4, Name = "Tablet", Price = 400.00m, Category = "Electronics" }
            };
            
            // Use LINQ to query products
            var expensiveElectronics = products
                .Where(p => p.Category == "Electronics" && p.Price > 500)
                .OrderBy(p => p.Price)
                .Select(p => new { p.Name, p.Price });
                
            Console.WriteLine("Expensive Electronics:");
            foreach (var item in expensiveElectronics)
            {
                Console.WriteLine($"{item.Name}: ${item.Price}");
            }
        }
    }
    
    class Product
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public decimal Price { get; set; }
        public string Category { get; set; }
    }
}`,
			"csharp",
		},
	}

	// Create a consistent style for all languages
	freeze := freezelib.New().
		WithTheme("github-dark").
		WithFont("JetBrains Mono", 14).
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(15, 0, 8).
		WithPadding(20)

	for _, lang := range languages {
		fmt.Printf("📝 Generating %s example...\n", lang.name)

		svgData, err := freeze.GenerateFromCode(lang.code, lang.lang)
		if err != nil {
			fmt.Printf("❌ Error with %s: %v\n", lang.name, err)
			continue
		}

		filename := fmt.Sprintf("output/language_%s.svg", lang.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Language comparison with the same algorithm
func languageComparisonExample() {
	fmt.Println("\n🔄 Same Algorithm in Different Languages")
	fmt.Println("---------------------------------------")

	// Fibonacci implementation in different languages
	fibImplementations := []struct {
		name string
		code string
		lang string
	}{
		{
			"go",
			`package main

import "fmt"

// Recursive Fibonacci implementation
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Iterative Fibonacci implementation
func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 10
	fmt.Printf("Fibonacci(%d) = %d (recursive)\n", n, fibonacci(n))
	fmt.Printf("Fibonacci(%d) = %d (iterative)\n", n, fibonacciIterative(n))
}`,
			"go",
		},
		{
			"python",
			`def fibonacci_recursive(n):
    """Calculate Fibonacci number recursively."""
    if n <= 1:
        return n
    return fibonacci_recursive(n-1) + fibonacci_recursive(n-2)

def fibonacci_iterative(n):
    """Calculate Fibonacci number iteratively."""
    if n <= 1:
        return n
    
    a, b = 0, 1
    for i in range(2, n+1):
        a, b = b, a + b
    return b

def fibonacci_dynamic(n):
    """Calculate Fibonacci number using dynamic programming."""
    memo = {0: 0, 1: 1}
    
    def fib(n):
        if n not in memo:
            memo[n] = fib(n-1) + fib(n-2)
        return memo[n]
    
    return fib(n)

# Test the functions
n = 10
print(f"Fibonacci({n}) = {fibonacci_recursive(n)} (recursive)")
print(f"Fibonacci({n}) = {fibonacci_iterative(n)} (iterative)")
print(f"Fibonacci({n}) = {fibonacci_dynamic(n)} (dynamic)")`,
			"python",
		},
		{
			"javascript",
			`// Recursive Fibonacci implementation
function fibonacciRecursive(n) {
    if (n <= 1) return n;
    return fibonacciRecursive(n - 1) + fibonacciRecursive(n - 2);
}

// Iterative Fibonacci implementation
function fibonacciIterative(n) {
    if (n <= 1) return n;
    
    let a = 0, b = 1;
    for (let i = 2; i <= n; i++) {
        const temp = a + b;
        a = b;
        b = temp;
    }
    return b;
}

// Fibonacci with memoization
function fibonacciMemoized(n, memo = {}) {
    if (n in memo) return memo[n];
    if (n <= 1) return n;
    
    memo[n] = fibonacciMemoized(n - 1, memo) + fibonacciMemoized(n - 2, memo);
    return memo[n];
}

// Test the functions
const n = 10;
console.log(Fibonacci(${n}) = ${fibonacciRecursive(n)} (recursive));
console.log(Fibonacci(${n}) = ${fibonacciIterative(n)} (iterative));
console.log(Fibonacci(${n}) = ${fibonacciMemoized(n)} (memoized));`,
			"javascript",
		},
		{
			"rust",
			`fn fibonacci_recursive(n: u32) -> u32 {
    match n {
        0 => 0,
        1 => 1,
        _ => fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2),
    }
}

fn fibonacci_iterative(n: u32) -> u32 {
    match n {
        0 => 0,
        1 => 1,
        _ => {
            let mut a = 0;
            let mut b = 1;
            
            for _ in 2..=n {
                let temp = a + b;
                a = b;
                b = temp;
            }
            
            b
        }
    }
}

fn main() {
    let n = 10;
    println!("Fibonacci({}) = {} (recursive)", n, fibonacci_recursive(n));
    println!("Fibonacci({}) = {} (iterative)", n, fibonacci_iterative(n));
}`,
			"rust",
		},
	}

	// Create a consistent style for comparison
	freeze := freezelib.New().
		WithTheme("dracula").
		WithFont("Fira Code", 14).
		WithWindow(true).
		WithLineNumbers(true).
		WithPadding(20)

	for _, impl := range fibImplementations {
		fmt.Printf("🧮 Generating Fibonacci in %s...\n", impl.name)

		svgData, err := freeze.GenerateFromCode(impl.code, impl.lang)
		if err != nil {
			fmt.Printf("❌ Error with %s: %v\n", impl.name, err)
			continue
		}

		filename := fmt.Sprintf("output/fibonacci_%s.svg", impl.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Multi-language project example
func multiLanguageProjectExample() {
	fmt.Println("\n🌐 Multi-language Project")
	fmt.Println("-------------------------")

	// Different files in a web project
	projectFiles := []struct {
		name     string
		code     string
		lang     string
		filename string
	}{
		{
			"HTML",
			`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Task Manager</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="container">
        <h1>Task Manager</h1>
        
        <div class="task-form">
            <input type="text" id="taskInput" placeholder="Add a new task...">
            <button id="addTask">Add</button>
        </div>
        
        <ul id="taskList" class="task-list"></ul>
        
        <div class="stats">
            <p>Total tasks: <span id="totalTasks">0</span></p>
            <p>Completed: <span id="completedTasks">0</span></p>
        </div>
    </div>
    
    <script src="app.js"></script>
</body>
</html>`,
			"html",
			"index.html",
		},
		{
			"CSS",
			`/* Task Manager Styles */
* {
    box-sizing: border-box;
    margin: 0;
    padding: 0;
}

body {
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
    line-height: 1.6;
    background-color: #f5f5f5;
    color: #333;
}

.container {
    max-width: 800px;
    margin: 2rem auto;
    padding: 2rem;
    background-color: white;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

h1 {
    text-align: center;
    margin-bottom: 2rem;
    color: #2c3e50;
}

.task-form {
    display: flex;
    margin-bottom: 1.5rem;
}

#taskInput {
    flex: 1;
    padding: 0.75rem;
    border: 1px solid #ddd;
    border-radius: 4px 0 0 4px;
    font-size: 1rem;
}

#addTask {
    padding: 0.75rem 1.5rem;
    background-color: #3498db;
    color: white;
    border: none;
    border-radius: 0 4px 4px 0;
    cursor: pointer;
    font-size: 1rem;
}

.task-list {
    list-style: none;
    margin-bottom: 1.5rem;
}

.task-list li {
    padding: 1rem;
    border-bottom: 1px solid #eee;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.task-list li.completed {
    text-decoration: line-through;
    color: #7f8c8d;
}

.stats {
    display: flex;
    justify-content: space-between;
    color: #7f8c8d;
    font-size: 0.9rem;
}`,
			"css",
			"styles.css",
		},
		{
			"JavaScript",
			`// Task Manager App
document.addEventListener('DOMContentLoaded', () => {
    // DOM Elements
    const taskInput = document.getElementById('taskInput');
    const addTaskBtn = document.getElementById('addTask');
    const taskList = document.getElementById('taskList');
    const totalTasksEl = document.getElementById('totalTasks');
    const completedTasksEl = document.getElementById('completedTasks');
    
    // Task array
    let tasks = JSON.parse(localStorage.getItem('tasks')) || [];
    
    // Initial render
    renderTasks();
    updateStats();
    
    // Event Listeners
    addTaskBtn.addEventListener('click', addTask);
    taskInput.addEventListener('keypress', e => {
        if (e.key === 'Enter') addTask();
    });
    
    // Add a new task
    function addTask() {
        const taskText = taskInput.value.trim();
        if (taskText === '') return;
        
        tasks.push({
            id: Date.now(),
            text: taskText,
            completed: false
        });
        
        saveToLocalStorage();
        renderTasks();
        updateStats();
        
        taskInput.value = '';
        taskInput.focus();
    }
    
    // Toggle task completion
    function toggleTask(id) {
        tasks = tasks.map(task => 
            task.id === id ? { ...task, completed: !task.completed } : task
        );
        
        saveToLocalStorage();
        renderTasks();
        updateStats();
    }
    
    // Delete a task
    function deleteTask(id) {
        tasks = tasks.filter(task => task.id !== id);
        
        saveToLocalStorage();
        renderTasks();
        updateStats();
    }
    
    // Render tasks to DOM
    function renderTasks() {
        taskList.innerHTML = '';
        
        tasks.forEach(task => {
            const li = document.createElement('li');
            li.className = task.completed ? 'completed' : '';
            
            li.innerHTML = 
                <span onclick="toggleTask(${task.id})">${task.text}</span>
                <button onclick="deleteTask(${task.id})">Delete</button>
            ;
            
            taskList.appendChild(li);
        });
    }
    
    // Update statistics
    function updateStats() {
        totalTasksEl.textContent = tasks.length;
        completedTasksEl.textContent = tasks.filter(task => task.completed).length;
    }
    
    // Save to localStorage
    function saveToLocalStorage() {
        localStorage.setItem('tasks', JSON.stringify(tasks));
    }
    
    // Expose functions to global scope for inline event handlers
    window.toggleTask = toggleTask;
    window.deleteTask = deleteTask;
});`,
			"javascript",
			"app.js",
		},
	}

	// Create a consistent style for the project files
	freeze := freezelib.New().
		WithTheme("github").
		WithFont("Cascadia Code", 13).
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(10, 0, 5).
		WithPadding(20)

	for _, file := range projectFiles {
		fmt.Printf("📄 Generating %s file (%s)...\n", file.name, file.filename)

		svgData, err := freeze.GenerateFromCode(file.code, file.lang)
		if err != nil {
			fmt.Printf("❌ Error with %s: %v\n", file.name, err)
			continue
		}

		filename := fmt.Sprintf("output/project_%s.svg", file.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Language-specific features example
func languageSpecificFeaturesExample() {
	fmt.Println("\n✨ Language-Specific Features")
	fmt.Println("----------------------------")

	// Language-specific code features
	features := []struct {
		name        string
		code        string
		lang        string
		description string
	}{
		{
			"go_concurrency",
			`package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	
	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Wait for workers to finish
	wg.Wait()
	close(results)
	
	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}`,
			"go",
			"Go Concurrency with Goroutines and Channels",
		},
		{
			"python_decorators",
			`import time
import functools
from typing import Callable, TypeVar, Any

T = TypeVar('T')

def timer(func: Callable[..., T]) -> Callable[..., T]:
    """Decorator that prints the execution time of a function."""
    @functools.wraps(func)
    def wrapper(*args: Any, **kwargs: Any) -> T:
        start_time = time.time()
        result = func(*args, **kwargs)
        end_time = time.time()
        print(f"{func.__name__} executed in {end_time - start_time:.4f} seconds")
        return result
    return wrapper

def memoize(func: Callable[..., T]) -> Callable[..., T]:
    """Decorator that caches function results."""
    cache = {}
    
    @functools.wraps(func)
    def wrapper(*args: Any, **kwargs: Any) -> T:
        key = str(args) + str(kwargs)
        if key not in cache:
            cache[key] = func(*args, **kwargs)
        return cache[key]
    return wrapper

@timer
@memoize
def fibonacci(n: int) -> int:
    """Calculate the nth Fibonacci number."""
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

# Test the decorated function
print(fibonacci(30))  # First call will be timed
print(fibonacci(30))  # Second call will use cached result`,
			"python",
			"Python Decorators and Type Annotations",
		},
		{
			"typescript_generics",
			`// TypeScript Generics and Interfaces
interface Repository<T> {
  getById(id: string): Promise<T>;
  getAll(): Promise<T[]>;
  create(item: T): Promise<T>;
  update(id: string, item: T): Promise<T>;
  delete(id: string): Promise<boolean>;
}

interface User {
  id?: string;
  name: string;
  email: string;
  role: 'admin' | 'user' | 'guest';
  createdAt?: Date;
}

class UserRepository implements Repository<User> {
  private users: Map<string, User> = new Map();
  
  async getById(id: string): Promise<User> {
    const user = this.users.get(id);
    if (!user) {
      throw new Error(User with id ${id} not found);
    }
    return user;
  }
  
  async getAll(): Promise<User[]> {
    return Array.from(this.users.values());
  }
  
  async create(user: User): Promise<User> {
    const id = Math.random().toString(36).substring(2, 9);
    const newUser = { 
      ...user, 
      id, 
      createdAt: new Date() 
    };
    this.users.set(id, newUser);
    return newUser;
  }
  
  async update(id: string, user: User): Promise<User> {
    if (!this.users.has(id)) {
      throw new Error('User with id ${id} not found');
    }
    const updatedUser = { ...user, id };
    this.users.set(id, updatedUser);
    return updatedUser;
  }
  
  async delete(id: string): Promise<boolean> {
    return this.users.delete(id);
  }
}

// Usage example
async function main() {
  const userRepo = new UserRepository();
  
  const newUser = await userRepo.create({
    name: 'John Doe',
    email: 'john@example.com',
    role: 'admin'
  });
  
  console.log('Created user:', newUser);
  
  const allUsers = await userRepo.getAll();
  console.log('All users:', allUsers);
}

main().catch(console.error);`,
			"typescript",
			"TypeScript Generics, Interfaces and Type Safety",
		},
	}

	// Create a consistent style for the feature examples
	freeze := freezelib.New().
		WithTheme("one-dark").
		WithFont("JetBrains Mono", 13).
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(15, 0, 8).
		WithPadding(20)

	for _, feature := range features {
		fmt.Printf("🔍 Generating %s example...\n", feature.name)

		svgData, err := freeze.GenerateFromCode(feature.code, feature.lang)
		if err != nil {
			fmt.Printf("❌ Error with %s: %v\n", feature.name, err)
			continue
		}

		filename := fmt.Sprintf("output/feature_%s.svg", feature.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s - %s\n", filename, feature.description)
	}
}
