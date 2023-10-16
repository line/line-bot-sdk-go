import os
import subprocess
import sys


def run_command(command):
    print(command)
    env = os.environ.copy()
    env["GO_POST_PROCESS_FILE"] = "./postprocess-file.sh"
    proc = subprocess.run(command, shell=True, text=True, capture_output=True, env=env)

    if len(proc.stdout) != 0:
        print("\n\nSTDOUT:\n\n")
        print(proc.stdout)

    if len(proc.stderr) != 0:
        print("\n\nSTDERR:\n\n")
        print(proc.stderr)
        print("\n\n")

    if proc.returncode != 0:
        print(f"\n\nCommand '{command}' returned non-zero exit status {proc.returncode}.")
        sys.exit(1)

    return proc.stdout.strip()


def read_files_from_openapi_generator(directory):
    file_path = os.path.join(directory, ".openapi-generator", "FILES")

    if not os.path.exists(file_path):
        print(f"Error: File {file_path} does not exist.")
        return []

    try:
        with open(file_path, 'r') as f:
            content = f.read()
    except Exception as e:
        print(f"Error reading file: {e}")
        return []

    file_list = content.strip().split("\n")
    return file_list


def generate_clients():
    components = [
        "shop.yml",
        "channel-access-token.yml",
        "insight.yml",
        "liff.yml",
        "manage-audience.yml",
        "module-attach.yml",
        "module.yml",
        "messaging-api.yml",
    ]

    for sourceYaml in components:
        output_path = 'linebot/' + sourceYaml.replace('.yml', '').replace('-', '_')

        orig_files = read_files_from_openapi_generator(output_path)

        command = f'''java \\
                    -cp ./tools/openapi-generator-cli.jar:./generator/target/line-bot-sdk-go-generator-openapi-generator-1.0.0.jar \\
                    org.openapitools.codegen.OpenAPIGenerator \\
                    generate \\
                    -g line-bot-sdk-go-generator \\
                    -e pebble \\
                    --model-package model \\
                    --api-package api \\
                    --package-name {sourceYaml.replace('.yml', '').replace('-', '_')} \\
                    --enable-post-process-file \\
                    -o {output_path.replace("-", "_")} \\
                    -i line-openapi/{sourceYaml} \\
                  '''
        run_command(command)

        curr_files = read_files_from_openapi_generator(output_path)

        for f in set(orig_files) - set(curr_files):
            run_command(f'rm -rf {f}')


def generate_webhook():
    source_yaml = "webhook.yml"
    output_path = 'linebot/webhook'

    orig_files = read_files_from_openapi_generator(output_path)

    command = f'''java \\
                    -cp ./tools/openapi-generator-cli.jar:./generator/target/line-bot-sdk-go-generator-openapi-generator-1.0.0.jar \\
                    org.openapitools.codegen.OpenAPIGenerator \\
                    generate \\
                    --global-property apiTest=false,modelDocs=false,apiDocs=false \\
                    --template-dir ./generator/src/main/resources \\
                    --enable-post-process-file \\
                    --additional-properties parse.go=true \\
                    -e pebble \\
                    --model-package model \\
                    --api-package api \\
                    --package-name webhook \\
                    -g line-bot-sdk-go-generator \\
                    -o {output_path} \\
                    -i line-openapi/{source_yaml} \\
                  '''
    run_command(command)

    curr_files = read_files_from_openapi_generator(output_path)

    for f in set(orig_files) - set(curr_files):
        run_command(f'rm -rf {f}')


def main():
    os.chdir(os.path.dirname(os.path.abspath(__file__)))

    os.chdir("generator")
    run_command('mvn package -DskipTests=true')
    os.chdir("..")

    generate_clients()
    generate_webhook()


if __name__ == "__main__":
    main()
