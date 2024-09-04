import os


def count_lines_of_code(directory):
    total_lines = 0
    for root, dirs, files in os.walk(directory):
        for file in files:
            if file.endswith(('.py', '.go', '.js', '.jsx', '.ts', '.tsx', '.html', '.css')):
                file_path = os.path.join(root, file)
                with open(file_path, 'r', encoding='utf-8') as f:
                    total_lines += sum(1 for line in f if line.strip())
    return total_lines


def main():
    server_dir = './server'
    client_dir = './client/src'

    server_lines = count_lines_of_code(server_dir)
    client_lines = count_lines_of_code(client_dir)

    print(f"Server code: {server_lines} lines")
    print(f"Client code: {client_lines} lines")
    print(f"Total lines of code: {server_lines + client_lines}")


if __name__ == "__main__":
    main()
