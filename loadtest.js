import http from 'k6/http';
import { check, sleep } from 'k6';
import { Rate } from 'k6/metrics';
import { Counter } from 'k6/metrics';

export let options = {
    stages: [
        { duration: '1m', target: 100 }, // ramp up to 100 users
        { duration: '5m', target: 100 }, // stay at 100 users
        { duration: '1m', target: 0 },   // ramp down to 0 users
    ],
    };
export let errorRate = new Rate('errors');
export let errorCount = new Counter('error_count');
export default function () {
    let url = 'http://localhost:8080';
    let params = {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9'
        },
    };

    let res = http.get(url, params);

    check(res, {
        'is status 200': (r) => r.status === 200,
        'response time < 200ms': (r) => r.timings.duration < 200,
    }) || errorRate.add(1);
    
    if (res.status !== 200) {
        errorCount.add(1);
    }

    sleep(1);
}