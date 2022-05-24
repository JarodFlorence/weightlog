<?php

namespace Danijwilliams\Dinero\Builders;

use Danijwilliams\Dinero\Models\Product;

class ProductBuilder extends Builder
{
    protected $entity = 'products';
    protected $model = Product::class;
}
