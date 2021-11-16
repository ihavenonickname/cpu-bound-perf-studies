import argparse
import json
import sys
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

    total_ms = 0

    print('Sending', args.requests, 'requests')
    print('Using', args.max_workers, 'max_workers')
    print('...')

    start = time()

    with ThreadPoolExecutor(max_workers=args.max_workers) as executor:
        futures = []

        for _ in range(args.requests):
            future = executor.submit(post_and_measure_time, args.url, payload)
            futures.append(future)

        for future in as_completed(futures):
            ms_elapsed = future.result()
            print('...request processed in', ms_elapsed, 'ms')
            total_ms += ms_elapsed

    end = time()
    total_elapsed = int((end - start) * 1000)

    print('Average per request:', int(total_ms / args.requests))
    print('Total elapsed:', total_elapsed)
    print('Total throuput:', int(total_elapsed / args.requests))

if __name__ == '__main__':
    main()
