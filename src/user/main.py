import json
import socket

SERVER_HOST = "127.0.0.1"
SERVER_PORT = 8080
FORMAT = "utf-8"

def __generate_request(raw_message:str):
    message_arr = raw_message.split(" ")
    request = {
        "command": message_arr[0],
        "name": message_arr[1],
        "size": int(message_arr[2])
    }
    return bytes(json.JSONEncoder().encode(request), FORMAT)

client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

if __name__ == "__main__":
    client.connect((SERVER_HOST, SERVER_PORT))
    while True:
        try:
            user_message = input("[CLIENT] Your message (\"EXIT\" to terminate): ")
            if user_message.upper() == "EXIT":
                print("[CLIENT] Terminated")
                client.close()
                break
            client.send(__generate_request(user_message))
        except KeyboardInterrupt:
            print("\n[CLIENT] Terminated")
            client.close()
            break