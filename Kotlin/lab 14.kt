data class Product(
    val id: Int,
    val name: String,
    val category: String,
    val price: Double,
    val discount: Double? = null,
    val description: String? = null
)

fun Product.finalPrice(): Double {
    val off = discount ?: 0.0
    return price - off
}

fun Product.details(): String {
    val desc = description ?: "No description"
    return "$id. $name ($category) - Rs.${finalPrice()} | $desc"
}

class CatalogueBuilder {
    var storeName: String = "My Store"
    val products = mutableListOf<Product>()

    fun product(
        id: Int,
        name: String,
        category: String,
        price: Double,
        discount: Double? = null,
        description: String? = null
    ) {
        products.add(Product(id, name, category, price, discount, description))
    }
}

fun catalogue(block: CatalogueBuilder.() -> Unit): CatalogueBuilder {
    val builder = CatalogueBuilder()
    builder.block()
    return builder
}

fun main() {
    val myStore = catalogue {
        storeName = "Campus Shop"
        product(1, "Mouse", "Electronics", 600.0, 50.0, "Wireless mouse")
        product(2, "Keyboard", "Electronics", 1200.0, null, "USB keyboard")
        product(3, "Notebook", "Stationery", 120.0, 20.0, null)
        product(4, "Pen Pack", "Stationery", 150.0)
        product(5, "Headphones", "Electronics", 1800.0, 300.0, "Wired headphones")
    }

    val allProducts = myStore.products

    println("Store: ${myStore.storeName}")
    println("\nAll Products:")
    for (p in allProducts) {
        println(p.details())
    }

    val budgetProducts = allProducts.filter { it.finalPrice() < 1000.0 }
    println("\nProducts under Rs.1000:")
    for (p in budgetProducts) {
        println("${p.name} - Rs.${p.finalPrice()}")
    }

    val prices = allProducts.map { it.finalPrice() }
    println("\nAll Final Prices: $prices")

    val totalAmount = prices.reduce { total, p -> total + p }
    println("\nTotal Catalogue Value: Rs.$totalAmount")

    val groupedByCategory = allProducts.groupBy { it.category }
    println("\nGrouped by Category:")
    for ((category, items) in groupedByCategory) {
        val names = items.map { it.name }
        println("$category: $names")
    }
}
