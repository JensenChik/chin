package tech.cuda.models

import org.jetbrains.exposed.sql.Database
import org.jetbrains.exposed.sql.transactions.transaction
import org.jetbrains.exposed.sql.SchemaUtils.create
import org.jetbrains.exposed.sql.SchemaUtils.drop

/**
 * Created by Jensen on 18-6-19.
 */

val tables = listOf(Group, Instance, Job, Machine, Operation, Task, User)

fun getDatabase(): Database {
    return Database.connect(
            user = "root",
            password = "root",
            url = "jdbc:mysql://localhost/chin",
            driver = "com.mysql.jdbc.Driver"
    )
}

fun createTables() {
    val db = getDatabase()
    transaction(db) {
        tables.forEach { create(it) }
    }
}

fun dropTables() {
    val db = getDatabase()
    transaction(db) {
        tables.forEach { drop(it) }
    }
}

fun rebuildTables() {
    dropTables()
    createTables()
}
