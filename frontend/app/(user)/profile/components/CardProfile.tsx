import React from 'react';
import Image from 'next/image';

interface CardProfileProps {
  imageUrl: string;
  name: string;
}

const CardProfile: React.FC<CardProfileProps> = ({ imageUrl, name }) => {
  const avatarUrl = imageUrl || '/avatar-male.svg';

  return (
    <div style={{ display: 'flex', alignItems: 'center' }}>
      <Image src={avatarUrl} alt={name} width={50} height={50} style={{ borderRadius: '50%', marginRight: '10px' }} />
      <span>{name}</span>
    </div>
  );
};

export default CardProfile;
