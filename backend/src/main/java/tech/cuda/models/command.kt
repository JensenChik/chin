package tech.cuda.models

import org.jetbrains.exposed.sql.Database
import org.jetbrains.exposed.sql.transactions.transaction
import org.jetbrains.exposed.sql.SchemaUtils
import tech.cuda.models.mappers.*

/**
 * Created by Jensen on 18-6-19.
 */

val tables = listOf(Groups, Instances, Jobs, Machines, Actions, Tasks, Users)

fun getDatabase(): Database {
    return Database.connect(
            user = "root",
            password = "qijinxiu",
            url = "jdbc:mysql://localhost/chin",
            driver = "com.mysql.jdbc.Driver"
    )
}

fun createTables() {
    val db = getDatabase()
    transaction(db) {
        SchemaUtils.create(Groups, Instances, Jobs, Machines, Actions, Tasks, Users)
    }
}

fun dropTables() {
    val db = getDatabase()
    transaction(db) {
        SchemaUtils.drop(Groups, Instances, Jobs, Machines, Actions, Tasks, Users)
    }
}

fun rebuildTables() {
    dropTables()
    createTables()
}

