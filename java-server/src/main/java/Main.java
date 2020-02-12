import io.grpc.Server;
import io.grpc.ServerBuilder;

import java.io.IOException;

public class Main {
    public static void main(String[] args) throws InterruptedException, IOException {
        Server server = ServerBuilder.forPort(9999).addService(new HelloServiceImpl()).build().start();
        server.awaitTermination();
    }
}
