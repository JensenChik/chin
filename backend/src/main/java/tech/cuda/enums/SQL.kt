package tech.cuda.enums


inline fun <reified T : Enum<T>> SQL(): String {
    return enumValues<T>().joinToString(",", prefix = "ENUM(", postfix = ")") { "'${it.name}'" }
}

