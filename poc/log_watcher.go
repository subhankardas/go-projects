package poc

// Log watcher library for Go. The log watcher will watch a file for changes and loads
// new lines to memory. We can also configure to load last few lines to memory initially
// when the watcher starts.
import (
	"bufio"
	"io"
	"os"
	"sync"
)

const (
	INITIAL_BUFF = 4
)

type logWatcher struct {
	file       *os.File
	fileReader *bufio.Reader
	loaded     bool
	buff       int8
}

// Create a new log watcher instance with the given filename.
func NewLogWatcher(filename string) *logWatcher {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	return &logWatcher{
		file:       file,
		fileReader: bufio.NewReader(file),
		loaded:     false,
		buff:       INITIAL_BUFF,
	}
}

// Start watching the log file. Initially only loads the last few lines as set by the buffer size.
func (w *logWatcher) Watch() {
	lines := make([]string, w.buff+1) // Initial buffer to store last few lines
	for {
		line, err := w.fileReader.ReadString('\n') // Read line
		if len(line) == 0 {                        // Avoid empty lines
			continue
		}

		if !w.loaded { // Initial buffer not loaded, load last few lines
			lines = append(lines, line)
			if len(lines) > int(w.buff) {
				lines = lines[len(lines)-int(w.buff):]
			}
			if err == io.EOF { // Reached end of file, load buffer
				w.loaded = true
				for _, line := range lines {
					w.Load(line)
				}
				continue
			}
			continue
		}

		w.Load(line) // Load line
	}
}

// Load a line into the watcher.
func (w *logWatcher) Load(line string) {
	print(line)
}

// Close the log watcher.
func (w *logWatcher) Close() {
	w.fileReader = nil
	w.file.Close()
}

// Entry point for testing the log watcher library.
func LogWatcherMain() {
	wg := sync.WaitGroup{}
	watcher := NewLogWatcher("..\\poc\\files\\log.txt")
	defer watcher.Close()

	wg.Add(1)
	go watcher.Watch() // Start log watcher in background

	wg.Wait()
}
