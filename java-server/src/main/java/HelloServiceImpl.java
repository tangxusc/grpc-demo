import com.test.hello.HelloGrpc;
import com.test.hello.HelloReply;
import com.test.hello.HelloRequest;
import io.grpc.stub.StreamObserver;

public class HelloServiceImpl extends HelloGrpc.HelloImplBase {

    @Override
    public void sayHello(HelloRequest request, StreamObserver<HelloReply> responseObserver) {
        System.out.println(request.getName());
        HelloReply build = HelloReply.newBuilder().setMessage("hello " + request.getName()).build();
        responseObserver.onNext(build);
        responseObserver.onCompleted();
    }
}
