<?php

namespace Danijwilliams\Dinero\Requests;

use Danijwilliams\Dinero\Builders\Builder;
use Danijwilliams\Dinero\Utils\RequestBuilder;

class DepositAccountRequestBuilder extends RequestBuilder
{
	public function __construct( Builder $builder ) {
		$this->parameters['fields'] = 'Name,AccountNumber';

		parent::__construct( $builder );
	}
}
