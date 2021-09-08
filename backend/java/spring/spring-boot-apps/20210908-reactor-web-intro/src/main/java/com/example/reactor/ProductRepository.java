package com.example.reactor;

import com.example.reactor.model.Product;
import org.reactivestreams.Publisher;
import org.springframework.data.mongodb.repository.ReactiveMongoRepository;
import reactor.core.publisher.Flux;

public interface ProductRepository
        extends ReactiveMongoRepository<Product, String> {
    Flux<Product> findByNameOrderByPrice(Publisher<String> name);
}
