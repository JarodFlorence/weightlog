<?php

namespace Danijwilliams\Dinero\Builders;

use Danijwilliams\Dinero\Models\Creditnote;

class CreditnoteBuilder extends Builder
{
    protected $entity = 'sales/creditnotes';
    protected $model = Creditnote::class;
}
