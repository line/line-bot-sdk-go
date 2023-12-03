package line.bot.generator;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.stream.Stream;

import org.junit.jupiter.api.Test;
import org.openapitools.codegen.ClientOptInput;
import org.openapitools.codegen.DefaultGenerator;
import org.openapitools.codegen.config.CodegenConfigurator;


/***
 * This test allows you to easily launch your code generation software under a debugger.
 * Then run this test under debug mode.  You will be able to step through your java code
 * and then see the results in the out directory.
 */
public class LineBotSdkNodejsGeneratorGeneratorTest {
    @Test
    public void messagingApi() throws IOException {
        generate("messaging-api");
    }

    @Test
    public void liff() throws IOException {
        generate("liff");
    }

    @Test
    public void webhook() throws IOException {
        generate("webhook");
    }

    private void generate(String target) throws IOException {
        Path outPath = Paths.get("out/" + target);
        if (outPath.toFile().exists()) {
            try (Stream<Path> stream = Files.walk(outPath)) {
                //noinspection ResultOfMethodCallIgnored
                stream.map(Path::toFile)
                        .forEach(File::delete);
            }
        }

        // to understand how the 'openapi-generator-cli' module is using 'CodegenConfigurator', have a look at the 'Generate' class:
        // https://github.com/OpenAPITools/openapi-generator/blob/master/modules/openapi-generator-cli/src/main/java/org/openapitools/codegen/cmd/Generate.java
        final CodegenConfigurator configurator = new CodegenConfigurator()
                .setTemplatingEngineName("pebble")
                .setTemplateDir("src/main/resources/line-bot-sdk-go-generator")
                .setGeneratorName("line-bot-sdk-go-generator") // use this codegen library
                .setModelPackage("model")
                .setApiPackage("api")
                .setPackageName(target.replace("-", "_"))
                .setInputSpec("../line-openapi/" + target + ".yml") // sample OpenAPI file
                .setOutputDir("out/" + target); // output directory

        final ClientOptInput clientOptInput = configurator.toClientOptInput();
        DefaultGenerator generator = new DefaultGenerator();
        generator.opts(clientOptInput).generate();
    }
}
