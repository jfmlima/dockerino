package adapter

import io.javalin.http.Context
import io.javalin.http.Handler
import org.eclipse.jetty.http.HttpStatus

class IdentityHandler : Handler {
    override fun handle(ctx: Context) {
        ctx.status(HttpStatus.OK_200).json(IdentityRepresenter(value = "Dummy"))
    }
}

data class IdentityRepresenter(val value: String)
