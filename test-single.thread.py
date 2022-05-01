import argparse
import json
import urllib3
from concurrent.futures import ThreadPoolExecutor, as_completed
from time import time

import requests

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

def post_and_measure_time(url, payload):
    start = time()

    requests.post(url, json=payload, verify=False)

    end = time()

    return int((end - start) * 1000)

def main():
    parser = argparse.ArgumentParser()

    parser.add_argument('--url')
    parser.add_argument('--max-workers', type=int)
    parser.add_argument('--requests', type=int)

    args = parser.parse_args()

    with open('payload.txt') as f:
        payload = json.load(f)

    print('Sending 1 request without measuring time')
    print('in order to warm up the JIT compiler...')

    post_and_measure_time(args.url, payload)

    print('...done.')

    print()

    print(f'Sending {args.requests} requests')
    print(f'Using {args.max_workers} max_workers')
    print('...')

    total_ms = 0

    start = time()

    with ThreadPoolExecutor(max_workers=args.max_workers) as executor:
        futures = []

        for _ in range(args.requests):
            future = executor.submit(post_and_measure_time, args.url, payload)
            futures.append(future)

        for i, future in enumerate(as_completed(futures), start=1):
            ms_elapsed = future.result()
            print(f'...request #{i:2d} processed in {ms_elapsed} ms')
            total_ms += ms_elapsed

    end = time()
    total_elapsed = int((end - start) * 1000)

    print(f'* Average per request: {int(total_ms / args.requests)} ms')
    print(f'* Total elapsed: {total_elapsed} ms')
    print(f'* Total throuput: {int(total_elapsed / args.requests)} ms')

if __name__ == '__main__':
    main()
