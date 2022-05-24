<?php

namespace Danijwilliams\Dinero\Builders;

use Danijwilliams\Dinero\Models\Contact;

class ContactBuilder extends Builder
{
    protected $entity = 'contacts';
    protected $model = Contact::class;
}
