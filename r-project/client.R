#' Example gRPC client
#'
#' Sends a message with a name and returns a message greeting the name.
#' @references \url{https://github.com/grpc/grpc/tree/master/examples/cpp/helloworld}

library(grpc)

spec <- system.file('books.proto', package = 'grpc')
impl <- read_services(spec)
host <- "internal-e0okt3nx.ew.gateway.dev:443"
client <- grpc_client(impl, host)

req <- client$ListBooks$build()
message <- client$ListBooks$call(req)

print(message)
print(as.list(message))