<?php namespace LasseRafn\Dinero\Builders;

use Danijwilliams\Dinero\Models\Payment;
use Danijwilliams\Dinero\Responses\ListResponse;
use Danijwilliams\Dinero\Utils\Request;

class PaymentBuilder extends Builder
{
	protected $entity         = 'payments';
	protected $model          = Payment::class;
	protected $collectionName = 'Payments';
	protected $responseClass  = ListResponse::class;

	public function __construct( Request $request, $entity ) {
		$this->entity = $entity;
		parent::__construct( $request );
	}
}
