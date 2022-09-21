import http from "k6/http";
import { check } from "k6";
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

const url = "http://127.0.0.1:3000/atomic/redis-lock-add";
const getValueUrl = "http://127.0.0.1:3000/atomic/value";
const resetValueUrl = "http://127.0.0.1:3000/atomic/reset-value";

export let options = {
  duration: "20s",
  preAllocatedVUs: 1000,
  VUs: 100,
  noVUConnectionReuse: true,
};
export function setup() {
  http.get(resetValueUrl);

  let res = http.get(getValueUrl);
  let body = JSON.parse(res.body);
  check(body, { "pre-check default value": (r) => r.redis_lock_value == 0 });
}

export default function () {
  let res = http.get(url);
  check(res, { "status ok: ": (r) => r.status == 200 });
}

export function teardown(data) {
  let res = http.get(getValueUrl);
  let body = JSON.parse(res.body);
  console.log("value of redis_lock_value: ", body.free_value);
  check(body, { "check result ": (r) => r.redis_lock_value == 100 });
}
