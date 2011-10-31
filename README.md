#The SOA Research Platform (SOAR)

1. SOA Complete: Prove benefits and capabilities from a minimal basis.
2. Extensions: Increase ease of use and efficiency by coordination, etc.
3. Management: Improve SOA life cycle by change control, etc.
4. Execution: Improve execution environment by working on QoS, QoE, reliability, dependability, scalability, etc.

The point here is *not* to make a platform to build production-level SOAs, but to create a platform which can be leveraged to try new ideas.

## Building
    export GOPATH=/path/to/checkout
    goinstall -clean frontend assets render

This will build all executables and install to /path/to/checkout/bin

## Example SOA

This design is meant to exercise a SOA. It is probably not the best way to do things.

This implementation is not complete.

### Document preparation (frontend)
This is the user interface. Document parts are uploaded, edited. Rendering is requested, and the rendered documents are acessed.

### Asset Service (asset)
Read and write files. This service stores the things needed to produce a document as well as the produced documents. It can be really simple, just using a path to represent the file and a bytestream to represent the contents.

#### Document Processing Service (render)
takes as input references to assets and the metadata needed to create a document. Gathers the assets and renders the document (using pandoc). The rendered document is placed on the asset server and a completion message is returned.

