## This repository demonstrates some Go concepts.

### 1. Synchronous package:
* **Max processes:** Controlling program run concurrency or parallel, by chance number of processes.
  
* **Mutex:** To ensure that only one Goroutine is running the critical section of code at any point of time to prevent race condition.
  
* **Once:** Run code only once time on first load.

* **Wait group:** To block Go routines and wait it executes completely.

### 2. Intf and Intf-extender package:
* **Extends interface:** To demo how to extends an interface in Go lang.

### 3. Channel package:
* **Buffered channel:** Use channel to communicate among Go routines.
  
* **Select:** To handle multi channel in one Go routine.

* **Unbuffered channel:** Unbuffered channel won't block until capacity channel full. 