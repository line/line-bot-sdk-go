package line.bot.generator.pebble;

import java.util.Collections;
import java.util.List;
import java.util.Map;

import org.openapitools.codegen.CodegenModel;

import io.pebbletemplates.pebble.extension.Function;
import io.pebbletemplates.pebble.template.EvaluationContext;
import io.pebbletemplates.pebble.template.PebbleTemplate;

public class HasDiscriminatorFunction implements Function {
    @Override
    public List<String> getArgumentNames() {
        return Collections.singletonList("model");
    }

    @Override
    public Object execute(Map<String, Object> args, PebbleTemplate self, EvaluationContext context, int lineNumber) {
        CodegenModel model = (CodegenModel) args.get("model");
        return model.allVars.stream()
                .anyMatch(it -> it.isDiscriminator);
    }
}
