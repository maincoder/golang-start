package main

import (
    "bufio"
    "errors"
    "fmt"
    "io"
    "os"
    "strings"
)

type vertex struct {
    inDegree  int
    outDegree int
    pagerank  float64
}

type edge struct {
    start int
    end   int
}

var vertexs []vertex
var edges []edge
var vertexID map[string]int = make(map[string]int)
var numVertex int = 0

func lineReader(filename string) (func() (string, error), error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    buf := bufio.NewReaderSize(f, 64)
    return func() (string, error) {
        line, isPrefix, err := buf.ReadLine()
        if err != nil {
            if err == io.EOF {
                if err := f.Close(); err != nil {
                    return "", err
                }
            }
            return "", err
        }
        if isPrefix {
            return "", errors.New("buffer size to small")
        }
        return string(line), nil
    }, nil
}

func addVertex(vertexName string) int {
    var ID int
    var ok bool
    if ID, ok = vertexID[vertexName]; !ok {
        ID = numVertex
        vertexID[vertexName] = ID
        vertexs = append(vertexs, vertex{})
        numVertex++
    }
    return ID
}

func read() {
    readline, err := lineReader("D:\\J2EE\\workspace\\helloGo\\src\\wt2g_inlinks.source")
    if err != nil {
        panic(err)
    }
    for {
        line, err := readline()
        if err != nil {
            if err == io.EOF {
                break
            }
            panic(err)
        }
        // Line format is like "ID1\tID2"
        sections := strings.Split(line, "\t")
        if len(sections) != 2 {
            panic(errors.New("Illegal line format"))
        }
        start := addVertex(sections[0])
        end := addVertex(sections[1])
        edges = append(edges, edge{start, end})
    }
}

func calcPagerank(alpha float64, numIterations int) {
    // Initialize out degree of every vertex
    for i := range edges {
        edge := &edges[i]
        vertexs[edge.start].outDegree++
        vertexs[edge.end].inDegree++
    }
    var I = make([]float64, numVertex)
    var S float64
    for i := 0; i < numVertex; i++ {
        vertexs[i].pagerank = 1 / float64(numVertex)
        I[i] = alpha / float64(numVertex)
    }
    // Calculate pagerank repeatedly until converge (numIterations times)
    for k := 0; k < numIterations; k++ {
        for i := range edges {
            edge := &edges[i]
            I[edge.end] += (1 - alpha) * vertexs[edge.start].pagerank / float64(vertexs[edge.start].outDegree)
        }
        S = 0
        for i := 0; i < numVertex; i++ {
            if vertexs[i].outDegree == 0 {
                S += (1 - alpha) * vertexs[i].pagerank / float64(numVertex)
            }
        }
        for i := 0; i < numVertex; i++ {
            vertexs[i].pagerank = I[i] + S
            I[i] = alpha / float64(numVertex)
        }
    }
}

func main() {
    read()
    calcPagerank(0.15, 30)
    fmt.Println("Done")
}