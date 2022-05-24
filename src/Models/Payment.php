<?php namespace Danijwilliams\Dinero\Models;

use Danijwilliams\Dinero\Utils\Model;

class Payment extends Model
{
    protected $entity     = 'payments';
    protected $primaryKey = 'Guid';

    public $Guid;
    public $DepositAccountNumber;
    public $ExternalReference;
    public $PaymentDate;
    public $Description;
    public $Amount;
    public $AmountInForeignCurrency;
}
