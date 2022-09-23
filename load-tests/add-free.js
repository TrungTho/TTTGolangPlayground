import http from "k6/http";
import { check } from "k6";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

const url = "http://127.0.0.1:3000/atomic/free-add";
const getValueUrl = "http://127.0.0.1:3000/atomic/value";
const resetValueUrl = "http://127.0.0.1:3000/atomic/reset-value";

export let options = {
  scenarios: {
    constant_request_rate: {
      executor: "constant-arrival-rate",
      rate: 1000,
      duration: "2s",
      preAllocatedVUs: 1000,
      maxVUs: 1000,
    },
  },
};
export function setup() {
  http.get(resetValueUrl);

  let res = http.get(getValueUrl);
  let body = JSON.parse(res.body);
  check(body, { "pre-check default value": (r) => r.free_value == 0 });
}

export default function () {
  let res = http.get(url);
  check(res, { "status ok: ": (r) => r.status == 200 });
}

export function teardown(data) {
  let res = http.get(getValueUrl);
  let body = JSON.parse(res.body);
  console.log("value of free_valiue: ", body.free_value);
  check(body, { "check result ": (r) => r.free_value == 100 });
}
