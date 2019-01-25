package tech.cuda.models

import org.jetbrains.exposed.sql.Database
import org.jetbrains.exposed.sql.transactions.transaction
import org.jetbrains.exposed.sql.SchemaUtils
import org.jetbrains.exposed.sql.Table

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
        tables.forEach {
            if (it is Table) {
                SchemaUtils.create(it)
            }
        }
    }
}

fun dropTables() {
    val db = getDatabase()
    transaction(db) {
        tables.forEach {
            if (it is Table) {
                SchemaUtils.drop(it)
            }
        }
    }
}

fun rebuildTables() {
    dropTables()
    createTables()
}
