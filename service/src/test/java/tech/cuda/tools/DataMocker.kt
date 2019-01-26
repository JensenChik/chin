package tech.cuda.tools

import org.jetbrains.exposed.sql.transactions.TransactionManager
import org.jetbrains.exposed.sql.transactions.transaction
import java.io.File

/**
 * Created by Jensen on 19-1-26.
 */

object DataMocker {
    fun loadMockedData() {
        val dataPath = "${File("").absolutePath}/src/test/resources/data"
        val tables = listOf("groups", "users")
        transaction {
            val mysql = TransactionManager.current()
            for (table in tables) {
                mysql.exec("""
                    load data local infile '$dataPath/$table.csv'
                    into table $table
                    fields terminated by ','
                    lines terminated by '\n'
                    ignore 1 lines
                """)
            }
        }
    }

}