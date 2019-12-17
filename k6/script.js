import http from "k6/http";
import { check } from "k6";

export let options = {
    thresholds: {
        http_req_duration: ["p(99)<100"],
    }
};

export default function() {
    http.get("http://cicd-sample-app.default.svc.cluster.local:8080/");
}
