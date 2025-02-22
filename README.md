<pre><code>   __  ___        __     _           
  /  |/  /__  ___/ /  __(_)__ _    __
 / /|_/ / _ \/ _  / |/ / / -_) |/|/ /
/_/  /_/\___/\_,_/|___/_/\__/|__,__/
</code></pre>

Transform your Go project's dependency graph into a dynamic, interactive visualization with modview.
This powerful tool takes the complexity out of your module graph, offering a clear and explorable view of your
project's dependencies.

![Modview Example Image](modview.png)

`modview` leverages the output of `go mod graph` to create a browser-based visualization, 
enabling you to navigate, search, and understand your dependency structure effortlessly. 
Whether you're optimizing your codebase, resolving version conflicts, or exploring the ecosystem 
surrounding your project, modview is your guide through the intricate web of Go modules.

## Features

**Interactive HTML Visualization:** Generate a dynamic, browser-friendly graph of your Go module dependencies.

**Dependency Highlighting:** Easily distinguish between picked and unpicked dependencies, as determined 
by the `Minimal Version Selection (MVS)` algorithm.

**Intuitive Navigation:** Zoom, pan, and explore large dependency graphs with ease.

**Search Functionality:** Quickly locate specific dependencies within your graph.

## Installation

### Option 1
To install modview, use the following command:

```bash
go install github.com/bayraktugrul/modview@latest
```
Ensure that your Go bin directory is in your system's PATH.

### Option 2

```bash
git clone https://github.com/bayraktugrul/modview.git
cd modview
go install
```

## Usage

Navigate to your Go project's root directory and run:
```
modview
```
This will generate a file named `dependency_tree.html` in the current directory. Open this file in a web browser to view
your module's dependency graph.

<img src="modview-opt.gif" alt="modview go mod graph dependency tree" style="width: 100%; max-width: 100%;" />

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)

## Contact

For questions and feedback, please open an issue on the GitHub repository.