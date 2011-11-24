# The SOA Research Platform (SOAR)

1. SOA Complete: Prove benefits and capabilities from a minimal basis.
2. Extensions: Increase ease of use and efficiency by coordination, etc.
3. Management: Improve SOA life cycle by change control, etc.
4. Execution: Improve execution environment by working on QoS, QoE, reliability, dependability, scalability, etc.

The point here is *not* to make a platform to build production-level SOAs, but to create a platform which can be leveraged to try new ideas.

This software is developed by a member of the [ICSY](http://www.icsy.de) research group.

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

## Service Invocation

Do something like:
    consumer, err := soar.NewConsumer(":1234")
    // handle the errs
    consumer.Invoke("method", arg1, arg2, ...)

The invoke API works almost as if you were calling the remote method directly, you just have to put "consumer.Invoke("method"" where you would normally put "method(".

The call is then encoded and tranmitted across the network. Right now there are two encoders: json and gob. As the interface for these handles both encoding and decoding (two-way messaging is needed), the interface is called Coder. The implementations are in jsoncoder and gobcoder. To use a Coder call something like:
	coder := gobcoder.NewCoder()
    consumer, err := soar.NewConsumerWithCoder(":1234", coder)

The default Coder is jsoncoder because I read the encoding. Both coders are light-weight wrappers around the encoders and decoders created for the rpc package. New Coders should be easy enought to make.

## Interaction style

The only interaction style supported now is synchronous request/response. This is because it is the easiest to deal with while debugging. It should be fairly easy to convert to an asynchronous form like is used in the rpc package after the core mechanic works (full capability invocation and response handling). I will not be going beyond request/response for now as anything more does not add functionality to SOA. I will get async working because that is a better invoke api which can more easily be adapted when transactions or coordination are added.

Some notes I made while creating the API:

soar defaults to request/response style call

if something else is used, then extra work should be needed to use it

for example, if coordination is used, then: consumer.InteractionMode = COORDINATION and change to the coordination API

Request/Response in a coordinated system simply sets the point the response should be sent to to this client and not another service. consumer.InvokeCoordinated("method", args, input, destination)

A transactional system would wrap a single invocation in a transaction. transaction := consumer.NewTransaction(); transaction.Invoke("method", args); transaction.Commit(); transaction.Abort()

to handle Async and Sync methods, the synchrous method could always wrap the async method. This would allow synchronous invocation on an asynchronous system.

The idea is to maintain the API so that services and their consumers do not have to be updated when a new feature is added.

Phase 1: write synchronous, request/reply; use to implement a "Complete SOA"