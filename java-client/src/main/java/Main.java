import com.test.hello.HelloGrpc;
import com.test.hello.HelloReply;
import com.test.hello.HelloRequest;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class Main {
    public static void main(String[] args) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 9999).usePlaintext().build();
        HelloGrpc.HelloBlockingStub helloBlockingStub = HelloGrpc.newBlockingStub(channel);
        HelloReply test1 = helloBlockingStub.sayHello(HelloRequest.newBuilder().setName("test1").build());
        System.out.println(test1.getMessage());
    }
}
