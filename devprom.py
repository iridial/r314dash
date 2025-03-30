import json
from http.server import BaseHTTPRequestHandler, HTTPServer
from urllib.parse import urlparse#, parse_qs
from time import sleep

# Define the request handler
class RequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        parsed_path = urlparse(self.path)
        if parsed_path.path == '/api/v1/query':
            self.handle_query(parsed_path.query)
        elif parsed_path.path == '/vpn/v1/publicip/ip':
            self.handle_vpn()
        elif parsed_path.path == '/api/v1/status':
            self.handle_status()
        else:
            self.send_response(404)
            self.end_headers()

    def handle_query(self, query):
        try:
            # Load data from the JSON file
            with open('devstats.json', 'r') as file:
                data = json.load(file)

            # Parse query parameters
            #query_params = parse_qs(query)
            #sleep(3)

            self.send_response(200)
            self.send_header('Content-Type', 'application/json')
            self.end_headers()
            self.wfile.write(json.dumps(data).encode('utf-8'))
        except Exception as e:
            self.send_response(500)
            self.end_headers()
            self.wfile.write(f'Error reading file: {str(e)}'.encode('utf-8'))

    def handle_vpn(self):
        try:
            # Load data from the JSON file
            with open('devvpn.json', 'r') as file:
                data = json.load(file)

            # Parse query parameters
            #query_params = parse_qs(query)
            #sleep(3)

            self.send_response(200)
            self.send_header('Content-Type', 'application/json')
            self.end_headers()
            self.wfile.write(json.dumps(data).encode('utf-8'))
        except Exception as e:
            self.send_response(500)
            self.end_headers()
            self.wfile.write(f'Error reading file: {str(e)}'.encode('utf-8'))

    def handle_status(self):
        status_response = {
            "status": "OK",
            "version": "1.0"
        }
        self.send_response(200)
        self.send_header('Content-Type', 'application/json')
        self.end_headers()
        self.wfile.write(json.dumps(status_response).encode('utf-8'))

# Set up the server
def run(server_class=HTTPServer, handler_class=RequestHandler, port=9000):
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    print(f'Starting server on port {port}...')
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        print("\nServer is shutting down...")
    finally:
        httpd.server_close()
        print("Server closed.")

# run in windows without output:
# python3 .\devprom.py *> $null
if __name__ == "__main__":
    run()
