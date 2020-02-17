# How to use protocol Buffers

## Entries in the ".proto" file
* Define a ".proto" file
* First you should define syntax. example: syntax="proto3"
* Because protocolBuffer is used to generate GO code, define these in a separate package
* command: “protoc -I=. --go_out=. <filename>.proto
  * The command does a few things:
    * -I specifies the input source directory.
    * --go_out indicates where the generated Go source files will go.
.   * <proto file> is the name of the file to generate the source from
