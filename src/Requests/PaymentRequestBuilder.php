<?php namespace Danijwilliams\Dinero\Requests;

use Danijwilliams\Dinero\Builders\Builder;
use Danijwilliams\Dinero\Utils\RequestBuilder;

class PaymentRequestBuilder extends RequestBuilder
{
	public function __construct( Builder $builder )
	{
		$this->parameters['fields'] = 'Guid,DepositAccountNumber,ExternalReference,PaymentDate,Description,Amount,AmountInForeignCurrency';

		parent::__construct( $builder );
	}
}
