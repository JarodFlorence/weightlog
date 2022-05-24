<?php

namespace Danijwilliams\Dinero\Builders;

use Danijwilliams\Dinero\Models\PurchaseVoucher;

class PurchaseVoucherBuilder extends Builder
{
    protected $entity = 'vouchers/purchase';
    protected $model = PurchaseVoucher::class;
}
