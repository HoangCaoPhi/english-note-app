import React from 'react';
import { Card } from 'antd';

interface WordGroupCardProps {
  id: string;
  name: string;
  description: string;
  imageUrl?: string;
  wordCount: number;
  onClick?: () => void;
  className?: string;
}

const WordGroupCard: React.FC<WordGroupCardProps> = ({
  name,
  description,
  imageUrl,
  wordCount,
  onClick,
  className
}) => {
  return (
    <Card
      hoverable
      onClick={onClick}
      cover={
        imageUrl && (
          <div className="h-48 overflow-hidden">
            <img
              alt={name}
              src={imageUrl}
              className="w-full h-full object-cover"
            />
          </div>
        )
      }
      className={`shadow-md hover:shadow-lg transition-shadow ${className}`}
    >
      <div className="space-y-2">
        <h3 className="text-lg font-semibold text-gray-800">{name}</h3>
        <p className="text-gray-600 text-sm line-clamp-2">{description}</p>
        <div className="text-sm text-gray-500">
          {wordCount} {wordCount === 1 ? 'word' : 'words'}
        </div>
      </div>
    </Card>
  );
};

export default WordGroupCard;
