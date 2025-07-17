package line.bot.generator;

import java.io.File;
import java.util.HashSet;
import java.util.Map;

import org.openapitools.codegen.CodegenModel;
import org.openapitools.codegen.CodegenProperty;
import org.openapitools.codegen.CodegenType;
import org.openapitools.codegen.SupportingFile;
import org.openapitools.codegen.languages.AbstractGoCodegen;
import org.openapitools.codegen.model.ModelMap;
import org.openapitools.codegen.model.ModelsMap;

// https://github.com/OpenAPITools/openapi-generator/blob/master/modules/openapi-generator/src/main/java/org/openapitools/codegen/languages/AbstractGoCodegen.java
public class LineBotSdkGoGeneratorGenerator extends AbstractGoCodegen {
    protected String outputTestFolder = "";
    protected String testFolder = "tests";

    /**
     * Configures the type of generator.
     *
     * @return the CodegenType for this generator
     * @see org.openapitools.codegen.CodegenType
     */
    public CodegenType getTag() {
        return CodegenType.OTHER;
    }

    /**
     * Configures a friendly name for the generator.  This will be used by the generator
     * to select the library with the -g flag.
     *
     * @return the friendly name for the generator
     */
    public String getName() {
        return "line-bot-sdk-go-generator";
    }

    public String getHelp() {
        return "Generates a line-bot-sdk-go-generator client library.";
    }

    public LineBotSdkGoGeneratorGenerator() {
        super();
        embeddedTemplateDir = templateDir = "line-bot-sdk-go-generator";
        apiTestTemplateFiles.put("line-bot-sdk-go-generator/api_test.pebble", "_test.go");
        modelTemplateFiles.remove("model.mustache");
        modelTemplateFiles.put("line-bot-sdk-go-generator/model.pebble", ".go");
        apiTemplateFiles.remove("api-single.mustache");
        apiTemplateFiles.put("line-bot-sdk-go-generator/api.pebble", ".go");
    }

    @Override
    public void processOpts() {
        super.processOpts();
        supportingFiles.clear();
    }

    @Override
    public Map<String, ModelsMap> postProcessAllModels(Map<String, ModelsMap> objs) {
        Map<String, ModelsMap> stringModelsMapMap = super.postProcessAllModels(objs);
        HashSet<String> discriminators = new HashSet<>();
        for (ModelsMap modelsMap : stringModelsMapMap.values()) {
            for (ModelMap modelMap : modelsMap.getModels()) {
                CodegenModel model = modelMap.getModel();
                if (model.getDiscriminator() != null) {
                    System.out.println("Discriminator found: " + model.name);
                    discriminators.add(model.name);
                }
            }
        }

        for (ModelsMap modelsMap : stringModelsMapMap.values()) {
            for (ModelMap modelMap : modelsMap.getModels()) {
                CodegenModel model = modelMap.getModel();

                for (CodegenProperty var : model.allVars) {
                    if (var.baseType.equals("array") && discriminators.contains(var.complexType)) {
                        // CallbackRequest contains `[]Event`.
                        // in this case, baseType=array, complexType=Event
                        var.getVendorExtensions().put("x-is-discriminator-array", true);
                        model.getVendorExtensions().put("x-has-discriminator", true);
                        System.out.println("UnmarshalJSON[]: " + model.name + " " + var.name + " " + var.baseType + " " + var.complexType);
                    } else if (var.baseType.equals("map") && discriminators.contains(var.complexType)) {
                        // TextMessageV2 contains `map[string]SubstitutionObject`.
                        // in this case, baseType=map, complexType=SubstitutionObject
                        // UnmarshalJSON is not needed for this case.
                        var.getVendorExtensions().put("x-is-discriminator-map", true);
                    } else if (discriminators.contains(var.dataType)) {
                        var.getVendorExtensions().put("x-is-discriminator", true);
                        model.getVendorExtensions().put("x-has-discriminator", true);
                        System.out.println("UnmarshalJSON: " + model.name + " " + var.name + " " + var.baseType + " " + var.complexType);
                    }

                    var.getVendorExtensions().put("x-type", type(model, var));
                }
            }
        }
        return stringModelsMapMap;
    }

    private String type(CodegenModel model, CodegenProperty var) {
        boolean isDiscriminator =
                var.vendorExtensions.get("x-is-discriminator-array") != null ||
                var.vendorExtensions.get("x-is-discriminator") != null ||
                var.vendorExtensions.get("x-is-discriminator-map") != null;

        if (var.isEnum) {
            return model.classname + var.datatypeWithEnum;
        } else if (isDiscriminator) {
            return var.datatypeWithEnum + "Interface";
        } else if (var.isModel) {
            return "*" + var.datatypeWithEnum;
        } else {
            return var.datatypeWithEnum;
        }
    }

    @Override
    public String apiTestFileFolder() {
        return (outputTestFolder + File.separator + testFolder + File.separator + apiPackage().replace('.', File.separatorChar)).replace('/', File.separatorChar);
    }

    @Override
    public void setOutputDir(String dir) {
        super.setOutputDir(dir);
        if (this.outputTestFolder.isEmpty()) {
            setOutputTestFolder(dir);
        }
    }

    public void setOutputTestFolder(String outputTestFolder) {
        this.outputTestFolder = outputTestFolder;
    }

    @Override
    public String apiFileFolder() {
        return outputFolder + File.separator;
    }

    @Override
    public String modelFileFolder() {
        return (outputFolder + File.separator).replace("/", File.separator);
    }
}
