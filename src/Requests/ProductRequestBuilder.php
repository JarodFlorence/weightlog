<?php

namespace Danijwilliams\Dinero\Requests;

use Danijwilliams\Dinero\Builders\Builder;
use Danijwilliams\Dinero\Utils\RequestBuilder;

class ProductRequestBuilder extends RequestBuilder
{
    public function __construct(Builder $builder)
    {
        $this->parameters['fields'] = 'Name,ProductGuid,Quantity,Unit';

        parent::__construct($builder);
    }

    /**
     * Free-text search.
     *
     * @param $query
     *
     * @return $this
     */
    public function search($query)
    {
        $this->parameters['freeTextSearch'] = $query;

        return $this;
    }
}
