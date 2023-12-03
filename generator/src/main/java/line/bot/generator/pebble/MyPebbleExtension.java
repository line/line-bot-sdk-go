package line.bot.generator.pebble;

import java.util.HashMap;
import java.util.Map;

import io.pebbletemplates.pebble.extension.AbstractExtension;
import io.pebbletemplates.pebble.extension.Function;

public class MyPebbleExtension extends AbstractExtension {
    @Override
    public Map<String, Function> getFunctions() {
        HashMap<String, Function> map = new HashMap<>();
        map.put("endpoint", new EndpointFunction());
        map.put("hasDiscriminator", new HasDiscriminatorFunction());
        map.put("isOmitEmpty", new IsOmitEmptyFunction());
        return map;
    }
}
