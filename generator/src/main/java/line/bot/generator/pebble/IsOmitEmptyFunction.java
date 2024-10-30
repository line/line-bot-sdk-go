package line.bot.generator.pebble;

import java.util.Collections;
import java.util.List;
import java.util.Map;

import org.openapitools.codegen.CodegenProperty;

import io.pebbletemplates.pebble.extension.Function;
import io.pebbletemplates.pebble.template.EvaluationContext;
import io.pebbletemplates.pebble.template.PebbleTemplate;

public class IsOmitEmptyFunction implements Function {
    @Override
    public List<String> getArgumentNames() {
        return Collections.singletonList("var");
    }

    @Override
    public Object execute(Map<String, Object> args, PebbleTemplate self, EvaluationContext context, int lineNumber) {
        CodegenProperty var = (CodegenProperty) args.get("var");
        if (var.required) {
            return false;
        }
        if (var.isEnum) {
            // In golang, we are defining enum values like this:
            /*
            type EnumType string
            const (
                EnumTypeValue1 EnumType = "value1"
                EnumTypeValue2 EnumType = "value2"
            )
             */
            // As a result, the default value is empty string.
            return true;
        }
        if (var.isString) {
            // such as imageBackgroundColor in ButtonsTemplate.
            return true;
        }
        if (var.isInteger && isGreaterThanZero(var.minimum)) {
            // For optional integer types with a minimum value greater than 0,
            // it's acceptable to treat 0 as a null value
            return true;
        }
        if (var.isPrimitiveType) {
            return false;
        }
        return true;
    }

    private boolean isGreaterThanZero(String val) {
        return val != null && Integer.parseInt(val) > 0;
    }
}
