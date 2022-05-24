<?php

namespace Danijwilliams\Dinero\Models;

use Danijwilliams\Dinero\Utils\Model;

class DepositAccount extends Model
{
    protected $entity = 'accounts/deposit';
    protected $primaryKey = 'AccountNumber';

    public $AccountNumber;
    public $Name;
}
