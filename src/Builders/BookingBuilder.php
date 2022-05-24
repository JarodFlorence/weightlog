<?php namespace Danijwilliams\Dinero\Builders;

use Danijwilliams\Dinero\Models\Book;
use Danijwilliams\Dinero\Responses\ListResponse;
use Danijwilliams\Dinero\Utils\Request;

class BookingBuilder extends Builder
{
	protected $entity         = 'book';
	protected $model          = Book::class;
	protected $collectionName = 'Book';
	protected $responseClass  = ListResponse::class;

	public function __construct( Request $request, $entity ) {
		$this->entity = $entity;
		parent::__construct( $request );
	}
}
