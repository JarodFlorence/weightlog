<?php

namespace Danijwilliams\Dinero\Requests;

use Danijwilliams\Dinero\Builders\Builder;
use Danijwilliams\Dinero\Utils\RequestBuilder;

class PurchaseVoucherRequestBuilder extends RequestBuilder
{
    public function __construct(Builder $builder)
    {
        $this->parameters['fields'] = 'Number,Guid,ContactGuid,VoucherDate,PaymentDate,Status,Timestamp,VoucherNumber,FileGuid,RegionKey,$DepositAccountNumber,ExternalReference';

        parent::__construct($builder);
    }
}
