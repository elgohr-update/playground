package com.example.ordersvc.order.model;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class OrderValue {

    private Order order;
    private double value;
}
