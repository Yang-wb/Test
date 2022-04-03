import com.tulin.ConsistentHashRouter;
import com.tulin.Node;

import java.util.Arrays;

/**
 * Created by IntelliJ IDEA
 * User: Ginger
 * Date: 2018/8/17
 * Time: 9:44 AM
 */
public class ConsistenHashMain implements Node {
    private String idc;
    private String ip;
    private int port;

    private static void goRoute(ConsistentHashRouter<ConsistenHashMain> consistentHashRouter , String ... requestIps) {
        for (String requestIp: requestIps) {
            System.out.println(requestIp + " is route to " + consistentHashRouter.routeNode(requestIp));
        }
    }

    @Override
    public String getKey() {
        return idc + '-' + ip + ":" + port;
    }

    @Override
    public String toString() {
        return getKey();
    }

    public ConsistenHashMain(String idc, String ip, int port) {
        this.idc = idc;
        this.ip = ip;
        this.port = port;
    }

    public static void main(String[] argv) {
        ConsistenHashMain node1 = new ConsistenHashMain("IDC1","127.0.0.1",8080);
        ConsistenHashMain node2 = new ConsistenHashMain("IDC1","127.0.0.1",8081);
        ConsistenHashMain node3 = new ConsistenHashMain("IDC1","127.0.0.1",8082);
        ConsistenHashMain node4 = new ConsistenHashMain("IDC1","127.0.0.1",8083);

        ConsistentHashRouter<ConsistenHashMain> consistentHashRouter = new ConsistentHashRouter<>(Arrays.asList(node1,node2,node3,node4),10);
        String requestIP1 = "192.168.0.1";
        String requestIP2 = "192.168.0.2";
        String requestIP3 = "192.168.0.3";
        String requestIP4 = "192.168.0.4";
        String requestIP5 = "192.168.0.5";

        goRoute(consistentHashRouter,
                requestIP1,
                requestIP2,
                requestIP3,
                requestIP4,
                requestIP5);

        consistentHashRouter.removeNode(node3);
        System.out.println("-------------remove node online " + node3.getKey() + "------------");
        goRoute(consistentHashRouter,
                requestIP1,
                requestIP2,
                requestIP3,
                requestIP4,
                requestIP5);
    }
}
