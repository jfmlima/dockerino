package adapter

import io.javalin.Javalin
import io.javalin.apibuilder.ApiBuilder.*

class Api {
    private val app: Javalin = Javalin
        .create { it.enableCorsForAllOrigins() }
        .routes {
            get("whoami", IdentityHandler())
        }
        .start(8080)
        .apply {
        }

    fun stop() {
        app.stop()
    }
}