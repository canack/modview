package internal

const (
	Template = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dependency Graph</title>
    <script src="https://d3js.org/d3.v7.min.js"></script>
    <style>
        body, html {
            margin: 0;
            padding: 0;
            width: 100%;
            height: 100%;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background-color: #f8f9fa;
            overflow: hidden;
        }
        #graph-container {
            width: 100%;
            height: 100%;
            background-image: 
                radial-gradient(#e9ecef 1px, transparent 1px),
                radial-gradient(#e9ecef 1px, transparent 1px);
            background-size: 20px 20px;
            background-position: 0 0, 10px 10px;
        }
        .node {
            cursor: pointer;
            transition: all 0.3s ease;
        }
        .link {
            stroke: #6c757d;
            stroke-opacity: 0.8;
            fill: none;
            transition: stroke 0.3s, stroke-width 0.3s;
        }
        .node text {
            fill: #495057;
            text-anchor: middle;
            dominant-baseline: middle;
            font-weight: 500;
            font-size: 12px;
        }
        .tooltip {
            position: absolute;
            background-color: rgba(255, 255, 255, 0.95);
            border: 1px solid #dee2e6;
            padding: 10px;
            border-radius: 4px;
            pointer-events: none;
            box-shadow: 0 4px 6px rgba(0,0,0,0.1);
            font-size: 14px;
            transition: all 0.2s ease;
        }
        .highlighted {
            stroke: #228be6;
            stroke-width: 2px;
        }
        .node-highlighted rect {
            stroke: #228be6;
            stroke-width: 2px;
        }
        .link-highlighted {
            stroke: #228be6;
            stroke-width: 4px;
            filter: drop-shadow(0 0 3px #228be6);
        }
        .link-unpicked-highlighted {
            stroke: #ffa94d;
            stroke-width: 4px;
            filter: drop-shadow(0 0 3px #ffa94d);
        }
        .node-highlighted-text {
            font-weight: bold;
            fill: #228be6;
        }
        .node-highlighted-bg {
            fill: #e7f5ff;
        }
        .node-picked-highlight rect {
            stroke: #1c7ed6;
            stroke-width: 2px;
            filter: drop-shadow(0 0 5px rgba(28, 126, 214, 0.5));
        }
        .node-unpicked-highlight rect {
            stroke: #ffa94d;
            stroke-width: 1.5px;
        }
        .node-hover rect {
            stroke: #228be6;
            stroke-width: 2px;
            filter: drop-shadow(0 0 5px rgba(34, 139, 230, 0.5));
        }
        .node-unpicked-hover rect {
            stroke: #ffa94d;
            stroke-width: 2px;
            filter: drop-shadow(0 0 5px rgba(255, 169, 77, 0.5));
        }
        #repo-info {
            position: fixed;
            top: 20px;
            left: 20px;
            background-color: rgba(255, 255, 255, 0.8);
            padding: 10px;
            border-radius: 5px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.2);
            font-size: 14px;
            color: #333;
            z-index: 1000;
        }
        #legend {
            position: fixed;
            top: 70px;
            left: 20px;
            background-color: rgba(255, 255, 255, 0.95);
            padding: 15px;
            border-radius: 8px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
            font-size: 14px;
            margin-bottom: 10px;
        }
        .legend-item {
            display: flex;
            align-items: center;
            margin-bottom: 10px;
        }
        .legend-color {
            width: 24px;
            height: 24px;
            margin-right: 12px;
            border-radius: 4px;
            border: 2px solid rgba(0,0,0,0.1);
        }
        #search-container {
            position: fixed;
            top: 190px;
            left: 20px;
            display: flex;
            align-items: center;
            background-color: rgba(255, 255, 255, 0.95);
            padding: 10px;
            border-radius: 8px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
        }
        #search-input {
            padding: 8px;
            border: 1px solid #ced4da;
            border-radius: 4px;
            font-size: 14px;
            width: 200px;
        }
        #search-icon {
            margin-left: 10px;
            color: #6c757d;
            font-size: 20px;
            cursor: pointer;
        }
        #search-results {
            position: fixed;
            top: 230px;
            left: 20px;
            background-color: rgba(255, 255, 255, 0.95);
            padding: 10px;
            border-radius: 8px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
            font-size: 14px;
            color: #333;
            display: none;
        }
        #zoom-controls {
            position: fixed;
            bottom: 20px;
            left: 20px;
            display: flex;
            background-color: rgba(255, 255, 255, 0.95);
            border-radius: 8px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        .zoom-button {
            padding: 10px 15px;
            font-size: 18px;
            background-color: #f8f9fa;
            border: none;
            cursor: pointer;
            transition: background-color 0.3s;
        }
        .zoom-button:hover {
            background-color: #e9ecef;
        }
        .zoom-button:active {
            background-color: #dee2e6;
        }
        #zoom-in {
            border-right: 1px solid #dee2e6;
        }
        .notification {
            position: fixed;
            bottom: 20px;
            right: 20px;
            background-color: rgba(0, 0, 139, 0);
            color: #42A5F5;
            padding: 15px 25px;
            border-radius: 8px;
            border: 2px solid #42A5F5;
            box-shadow: 0 4px 12px rgba(0,0,0,0.2);
            font-size: 16px;
            opacity: 0;
            transform: translateY(20px);
            transition: opacity 0.5s ease, transform 0.5s ease;
            z-index: 1000;
        }
        .notification.show {
            opacity: 1;
            transform: translateY(0);
        }
    </style>
</head>
<body>
    <div id="repo-info">
        <a href="https://github.com/bayraktugrul/modview" target="_blank" style="text-decoration: none; color: #007bff;">Repository: bayraktugrul/modview</a>
    </div>
    <div id="legend">
        <div class="legend-item">
            <div class="legend-color" style="background-color: #e7f5ff; border-color: #1c7ed6;"></div>
            <span>Picked dependency by MVS algorithm</span>
        </div>
        <div class="legend-item">
            <div class="legend-color" style="background-color: #fff4e6; border-color: #ffa94d;"></div>
            <span>Unpicked dependency</span>
        </div>
    </div>
    <div id="search-container">
        <input type="text" id="search-input" placeholder="Search dependency...">
        <span id="search-icon">⏎</span>
    </div>
    <div id="search-results"></div>
    <div id="graph-container"></div>
    <div id="zoom-controls">
        <button id="zoom-in" class="zoom-button">+</button>
        <button id="zoom-out" class="zoom-button">-</button>
    </div>
    <div class="notification" id="notification">Copied!</div>
    <script>
        const data = {
            nodes: [
                {{- range $node := .Nodes }}
                { id: "{{ $node }}", picked: {{ if in $.MvsPicked $node }}true{{ else if in $.MvsUnpicked $node }}false{{ else }}null{{ end }} },
                {{- end }}
            ],
            links: [
                {{- range $edge := .Edges }}
                { source: "{{ $edge.From }}", target: "{{ $edge.To }}" },
                {{- end }}
            ],
            root: "{{ .Root }}"
        };

        const width = window.innerWidth;
        const height = window.innerHeight;

        const svg = d3.select("#graph-container")
            .append("svg")
            .attr("width", width)
            .attr("height", height);

        const g = svg.append("g");

        const hierarchy = d3.stratify()
            .id(d => d.id)
            .parentId(d => {
                const parent = data.links.find(link => link.target === d.id);
                return parent ? parent.source : null;
            })(data.nodes);

        const nodeCount = data.nodes.length;
        const dynamicWidth = Math.max(width, nodeCount * 80);
        const dynamicHeight = Math.max(height, nodeCount * 40);

        const treeLayout = d3.tree()
            .size([dynamicWidth - 200, dynamicHeight - 200])
            .separation((a, b) => (a.parent == b.parent ? 1 : 2) / (nodeCount > 50 ? 2 : 1));

        const treeData = treeLayout(hierarchy);

        const minVerticalSpacing = nodeCount > 50 ? 30 : 50;
        treeData.each(d => {
            d.y = height / 6 + d.depth * Math.max(80, minVerticalSpacing);
        });

        const zoom = d3.zoom()
            .scaleExtent([0.1, 4])
            .on("zoom", (event) => {
                g.attr("transform", event.transform);
            });

        svg.call(zoom);

        d3.select("#zoom-in").on("click", () => {
            svg.transition().duration(200).call(zoom.scaleBy, 1.5);
        });

        d3.select("#zoom-out").on("click", () => {
            svg.transition().duration(200).call(zoom.scaleBy, 0.5);
        });

        const link = g.selectAll(".link")
            .data(treeData.links())
            .enter().append("path")
            .attr("class", "link")
            .attr("d", d3.linkHorizontal()
                .x(d => d.y)
                .y(d => d.x))
            .style("stroke-width", nodeCount > 50 ? "0.5px" : (nodeCount > 20 ? "1px" : "2px"));

        const node = g.selectAll(".node")
            .data(treeData.descendants())
            .enter().append("g")
            .attr("class", "node")
            .attr("transform", d => "translate(" + d.y + "," + d.x + ")");

        const fontSize = nodeCount <= 15 ? 12 : (nodeCount > 50 ? 6 : (nodeCount > 20 ? 8 : 10));
        const nodeWidth = d => {
            if (nodeCount <= 15) {
                return Math.max(...treeData.descendants().map(n => n.data.id.length * (fontSize * 0.6) + 20));
            }
            return Math.max(Math.min(d.data.id.length * (fontSize * 0.6) + 20, 150), 80);
        };
        const nodeHeight = nodeCount <= 15 ? fontSize * 4 : fontSize * 3;

        node.append("rect")
            .attr("width", d => nodeWidth(d))
            .attr("height", nodeHeight)
            .attr("x", d => -nodeWidth(d) / 2)
            .attr("y", -nodeHeight / 2)
            .attr("rx", 6)
            .attr("ry", 6)
            .attr("fill", d => {
                if (d.data.id === data.root) return "#bbdefb";
                if (d.data.picked === true) return "#e7f5ff";
                if (d.data.picked === false) return "#fff4e6";
                return "#f5f5f5";
            })
            .attr("stroke", d => {
                if (d.data.id === data.root) return "#1e88e5";
                if (d.data.picked === true) return "#1c7ed6";
                if (d.data.picked === false) return "#ffa94d";
                return "#bdbdbd";
            })
            .attr("stroke-width", d => {
                if (d.data.picked === true) return 2;
                if (d.data.picked === false) return 1.5;
                return 2;
            });

        const nodeText = node.append("text")
            .attr("dy", "0.35em")
            .style("font-size", fontSize + "px")
            .attr("text-anchor", "middle")
            .attr("dominant-baseline", "middle");

        function truncateText(text, maxLength) {
            return text.length > maxLength ? text.slice(0, maxLength - 3) + "..." : text;
        }

        const tooltip = d3.select("body").append("div")
            .attr("class", "tooltip")
            .style("opacity", 0);

        nodeText.each(function(d) {
            const text = d3.select(this);
            const words = d.data.id.split(/(?=[A-Z@])/g);
            const lineHeight = 1.1;
            const y = text.attr("y");
            const dy = parseFloat(text.attr("dy"));
            let tspan = text.text(null).append("tspan").attr("x", 0).attr("y", y).attr("dy", dy + "em");
            let lineNumber = 0;
            let line = [];
            let isTruncated = false;

            words.forEach((word, i) => {
                line.push(word);
                tspan.text(line.join(""));
                if (tspan.node().getComputedTextLength() > nodeWidth(d) - 4) {
                    if (lineNumber === 1) {
                        line.pop();
                        tspan.text(truncateText(line.join(""), line.join("").length - 3));
                        isTruncated = true;
                        return;
                    }
                    line.pop();
                    tspan.text(line.join(""));
                    line = [word];
                    tspan = text.append("tspan").attr("x", 0).attr("y", y).attr("dy", ++lineNumber * lineHeight + dy + "em").text(word);
                }
            });

            if (isTruncated) {
                d3.select(this.parentNode)
                    .on("mouseover", function(event, d) {
                        tooltip.transition()
                            .duration(200)
                            .style("opacity", .9);
                        tooltip.html(d.data.id)
                            .style("left", (event.pageX + 10) + "px")
                            .style("top", (event.pageY - 28) + "px");
                    })
                    .on("mouseout", function(d) {
                        tooltip.transition()
                            .duration(500)
                            .style("opacity", 0);
                    });
            }
        });

        node.filter(d => d.data.id === data.root)
            .select("text")
            .attr("text-anchor", "middle")
            .attr("dominant-baseline", "middle")
            .selectAll("tspan")
            .attr("x", 0)
            .attr("dy", (d, i) => i ? "1.1em" : 0);

        const simulation = d3.forceSimulation(treeData.descendants())
            .force("collide", d3.forceCollide().radius(d => nodeWidth(d) / 2 + 10))
            .force("x", d3.forceX(d => d.y).strength(1))
            .force("y", d3.forceY(d => d.x).strength(1))
            .stop();

        for (let i = 0; i < 100; i++) {
            simulation.tick();
        }

        node.attr("transform", d => "translate(" + d.y + "," + d.x + ")");

        link.attr("d", d3.linkHorizontal()
            .x(d => d.y)
            .y(d => d.x));

        const rootNode = treeData.descendants()[0];
        const scale = nodeCount <= 15 ? 0.9 : 0.8;
        const initialX = width / 2 - rootNode.y * scale;
        const initialY = height / 6;
        svg.call(zoom.transform, d3.zoomIdentity.translate(initialX, initialY).scale(scale));

        node.on("mouseover", function(event, d) {
            highlightPathToRoot(d);
            highlightNodeAndRoot(d);
            if (d.data.picked === false) {
                d3.select(this).classed("node-unpicked-hover", true);
            } else {
                d3.select(this).classed("node-hover", true);
            }
        }).on("mouseout", function() {
            clearHighlight();
            clearNodeAndRootHighlight();
            d3.select(this).classed("node-hover", false);
            d3.select(this).classed("node-unpicked-hover", false);
        }).on("click", function(event, d) {
            const textToCopy = d.data.id;
            navigator.clipboard.writeText(textToCopy).then(() => {
                const notification = document.getElementById("notification");
                notification.classList.add("show");
                setTimeout(() => {
                    notification.classList.remove("show");
                }, 1000);
            });
            focusOnParentNode(d);
        });

        link.on("mouseover", function(event, d) {
            if (d.target.data.picked === false) {
                d3.select(this).classed("link-unpicked-highlighted", true);
            } else {
                d3.select(this).classed("link-highlighted", true);
            }
            highlightConnectedNodes(d);
        }).on("mouseout", function() {
            d3.select(this).classed("link-highlighted", false);
            d3.select(this).classed("link-unpicked-highlighted", false);
            clearNodeHighlight();
        });

        function highlightPathToRoot(node) {
            let current = node;
            while (current) {
                d3.select(current.node).select("rect").classed("highlighted", true);
                if (current.parent) {
                    const linkToParent = link.filter(l => l.target === current && l.source === current.parent);
                    linkToParent.classed("highlighted", true).style("stroke-width", "4px");
                }
                current = current.parent;
            }
        }

        function clearHighlight() {
            node.select("rect").classed("highlighted", false);
            link.classed("highlighted", false).style("stroke-width", null);
        }

        function highlightConnectedNodes(d) {
            d3.select(d.source.node).classed("node-highlighted", true);
            d3.select(d.target.node).classed("node-highlighted", true);
            d3.select(d.source.node).select("rect").classed("node-highlighted-bg", true);
            d3.select(d.target.node).select("rect").classed("node-highlighted-bg", true);
            d3.select(d.source.node).select("text").classed("node-highlighted-text", true);
            d3.select(d.target.node).select("text").classed("node-highlighted-text", true);
        }

        function clearNodeHighlight() {
            node.classed("node-highlighted", false);
            node.select("rect").classed("node-highlighted-bg", false);
            node.select("text").classed("node-highlighted-text", false);
        }

        function focusOnParentNode(d) {
            if (d.parent) {
                const scale = 2;
                const x = width / 2 - d.parent.y * scale;
                const y = height / 2 - d.parent.x * scale;

                svg.transition()
                    .duration(200)
                    .call(
                        zoom.transform,
                        d3.zoomIdentity.translate(x, y).scale(scale)
                    );
            }
        }

        function highlightNodeAndRoot(d) {
            const rootNode = treeData.descendants().find(node => node.data.id === data.root);
            const highlightClass = d.data.picked ? "node-picked-highlight" : "node-unpicked-highlight";
            
            d3.select(d.node).classed(highlightClass, true);
            d3.select(rootNode.node).classed(highlightClass, true);
        }

        function clearNodeAndRootHighlight() {
            node.classed("node-picked-highlight", false);
            node.classed("node-unpicked-highlight", false);
        }

        let searchResults = [];
        let currentSearchIndex = 0;

        function searchNodes(searchTerm) {
            searchResults = treeData.descendants().filter(node => node.data.id.toLowerCase().includes(searchTerm));
            currentSearchIndex = 0;
            if (searchResults.length > 0) {
                focusOnSearchResult();
                displaySearchResultsCount(searchResults.length);
            } else {
                hideSearchResultsCount();
            }
        }

        function focusOnSearchResult() {
            const matchedNode = searchResults[currentSearchIndex];
            if (matchedNode) {
                const scale = 2;
                const x = width / 2 - matchedNode.y * scale;
                const y = height / 2 - matchedNode.x * scale;

                svg.transition()
                    .duration(200)
                    .call(
                        zoom.transform,
                        d3.zoomIdentity.translate(x, y).scale(scale)
                    );

                node.classed("node-highlighted", d => d === matchedNode);
            }
        }

        function displaySearchResultsCount(count) {
            const searchResultsDiv = d3.select("#search-results");
            searchResultsDiv.style("display", "block");
            searchResultsDiv.html("Found " + count + " node" + (count > 1 ? "s" : ""));
        }

        function hideSearchResultsCount() {
            const searchResultsDiv = d3.select("#search-results");
            searchResultsDiv.style("display", "none");
        }

        d3.select("#search-icon").on("click", () => {
            const searchTerm = d3.select("#search-input").property("value").toLowerCase();
            if (searchTerm) {
                searchNodes(searchTerm);
            } else {
                hideSearchResultsCount();
            }
        });

        d3.select("#search-input").on("keypress", (event) => {
            if (event.key === "Enter") {
                if (searchResults.length > 0) {
                    currentSearchIndex = (currentSearchIndex + 1) % searchResults.length;
                    focusOnSearchResult();
                } else {
                    d3.select("#search-icon").dispatch("click");
                }
            }
        });
    </script>
</body>
</html>
`
)
