<?php

namespace Danijwilliams\Dinero\Builders;

use Danijwilliams\Dinero\Models\Invoice;

class InvoiceBuilder extends Builder
{
    protected $entity = 'invoices';
    protected $model = Invoice::class;
}
