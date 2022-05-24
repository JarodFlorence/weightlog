<?php

namespace Danijwilliams\Dinero\Models;

use Danijwilliams\Dinero\Utils\Model;

class Product extends Model
{
    protected $entity = 'products';
    protected $primaryKey = 'ProductGuid';

    public $CreatedAt;
    public $UpdatedAt;
    public $DeletedAt;
    public $ProductGuid;
    public $ProductNumber;
    public $Name;
    public $BaseAmountValue;
    public $Quantity;
    public $AccountNumber;
    public $Unit;
}
