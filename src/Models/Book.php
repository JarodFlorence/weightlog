<?php namespace Danijwilliams\Dinero\Models;

use Danijwilliams\Dinero\Utils\Model;

class Book extends Model
{
    protected $entity     = 'book';
    protected $primaryKey = 'Guid';

    public $Guid;
    public $Number;
    public $Timestamp;
}
