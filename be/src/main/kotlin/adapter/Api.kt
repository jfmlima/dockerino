package adapter

import io.javalin.Javalin
import io.javalin.apibuilder.ApiBuilder.get
import io.javalin.apibuilder.ApiBuilder.path

class Api {
    private val app: Javalin = Javalin
        .create { it.enableCorsForAllOrigins() }
        .routes {
            path("api") {
                get("whoami", IdentityHandler())
            }
        }
        .start(8080)
        .apply {}

    fun stop() {
        app.stop()
    }
}